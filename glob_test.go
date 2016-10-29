package main

import (
	"path/filepath"
	"testing"

	gglob "github.com/gobwas/glob"
	goglob "github.com/ryanuber/go-glob"
	v23glob "v.io/v23/glob"
)

type matcher func(s string) bool

func BenchmarkGlob(b *testing.B) {
	glob := "**.bad-request"
	match := []string{".bad-request", "svc.bad-request", "caller.dest.bad-request", "this.is.a.really.long.long.thing.that.will.end.up.matching.the.bad-request"}
	noMatch := []string{"", "bad-request", ".bad-request.", "random", "this.is.a.really.long.long.thing.that.will.not.end.up.matching.the.bad.request"}

	v23g, err := v23glob.Parse(glob)
	if err != nil {
		panic(err)
	}

	gg, err := gglob.Compile(glob, '.')
	if err != nil {
		panic(err)
	}

	runs := []struct {
		name string
		m    matcher
	}{
		{
			name: "go-glob",
			m: func(s string) bool {
				return goglob.Glob(glob, s)
			},
		},
		{
			name: "v23glob",
			m: func(s string) bool {
				return v23g.Head().Match(s)
			},
		},
		{
			name: "gobwas/glob",
			m: func(s string) bool {
				return gg.Match(s)
			},
		},
		{
			name: "filepath/match",
			m: func(s string) bool {
				ok, _ := filepath.Match(glob, s)
				return ok
			},
		},
	}

	for _, r := range runs {
		b.Run(r.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, m := range match {
					if !r.m(m) {
						panic("fail")
					}
				}
				for _, m := range noMatch {
					if r.m(m) {
						panic("fail")
					}
				}
			}
		})
	}
}
