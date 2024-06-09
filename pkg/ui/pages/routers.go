package pages

import (
	"bytes"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/silveiralexf/goflat/pkg/ui/elements"
	"github.com/silveiralexf/goflat/web"

	g "github.com/maragudk/gomponents"
)

func GetHandlerSet(timeout time.Duration) http.Handler {
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	slogger = slogger.With("env", "development")
	slog.SetDefault(slogger)

	r := chi.NewRouter()
	r.Use(middleware.Timeout(timeout))
	r.Use(middleware.Timeout(time.Millisecond * 600))
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	pathPrefix := os.Getenv("GOFLAT_ASSETS_PREFIX")
	if pathPrefix == "" {
		pathPrefix = filepath.Clean("dist")
	}

	uiFS, err := web.LoadEmbedFS(pathPrefix)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	r.Get("/assets", getAssetsHandler(uiFS))

	r.Get("/*", getPageHandler(indexPage()))
	r.Get("/about", getPageHandler(aboutPage()))
	r.Get("/contact", getPageHandler(contactPage()))
	r.Get("/photos", getPageHandler(photosPage()))
	r.Post("/photos", getPageHandler(photosPage()))
	return r
}

func getPageHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = elements.Page(title, r.URL.Path, body).Render(w)
	}
}

func getAssetsHandler(uiFS fs.FS) http.HandlerFunc {
	uiHandler := http.ServeMux{}
	err := fs.WalkDir(uiFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		b, err := fs.ReadFile(uiFS, path)
		if err != nil {
			return fmt.Errorf("failed to read ui file %s: %w", path, err)
		}

		fi, err := d.Info()
		if err != nil {
			return fmt.Errorf("failed to receive file info %s: %w", path, err)
		}

		paths := []string{fmt.Sprintf("/%s", path)}

		if paths[0] == "/index.html" {
			paths = append(paths, "/")
		}

		for _, path := range paths {
			uiHandler.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				http.ServeContent(w, r, d.Name(), fi.ModTime(), bytes.NewReader(b))
			})
		}
		return nil
	})
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return uiHandler.ServeHTTP
}
