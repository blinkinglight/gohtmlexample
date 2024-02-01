package html

import (
	"embed"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

//go:embed templates/*.html
var templates embed.FS

type HTML struct {
	viewData     map[string]interface{}
	httpRequest  *http.Request
	httpResponse http.ResponseWriter
}

type TemplateRender struct {
	template *template.Template
	data     map[string]interface{}
	html     *HTML
}

func (t *TemplateRender) Render(el string) {
	err := t.template.ExecuteTemplate(t.html.httpResponse, el, t.data)
	if err != nil {
		log.Printf("%+v", err)
	}
}

func New(w http.ResponseWriter, r *http.Request) *HTML {
	return &HTML{
		viewData:     get(r),
		httpRequest:  r,
		httpResponse: w,
	}
}

func (h *HTML) With(name string, data map[string]interface{}) *TemplateRender {
	templs := []string{"templates/components.html", "templates/" + name, "templates/index.html"}

	tpl, err := template.New(name).Funcs(h.newFuncMap()).ParseFS(templates, templs...)
	if err != nil {
		log.Printf("%+v", err)
	}

	vd := h.viewData
	for k, v := range data {
		vd[k] = v
	}

	return &TemplateRender{
		template: template.Must(tpl, err),
		data:     vd,
		html:     h,
	}
}

func (h *HTML) newFuncMap() template.FuncMap {
	return template.FuncMap{
		"RANDOM": func() uint32 {
			return uint32(rand.Int31n(999999999))
		},
		"LANG": func() string {
			return h.viewData["language"].(string)
		},
		"T": func(key string) string {
			return "translate:this:key:" + key
		},
	}
}

func get(r *http.Request) map[string]interface{} {
	i, ok := r.Context().Value("view-data").(map[string]interface{})
	if ok {
		return i
	}
	return make(map[string]interface{})
}
