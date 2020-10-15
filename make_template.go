package main

const makeTmpl = `
CC=g++
CCFLAGS = -std=c++11

{{.ExName}}: {{range .SFiles}}{{.}}.o {{end}}{{range .HFiles}}{{.}}.h {{end}}
	$(CC) $(CCFLAGS) {{range .SFiles}}{{.}}.o {{end}}-o {{.ExName}}

{{ $sources := .SFiles }}{{ $headers := .HFiles }}
{{range $s := $sources}}{{.}}.o: {{.}}.cpp {{range $h := $headers}}{{.}}.h {{end}}
	$(CC) $(CCFLAGS) {{.}}.cpp -c

{{end}}
clean:
	@rm *.o
	@rm {{.ExName}}
`
