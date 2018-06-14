package main

import "github.com/hunterlong/statup/plugin"

func OnHit(s *Service) {
	for _, p := range allPlugins {
		p.OnHit(s.ToP())
	}
}

func OnFailure(s *Service) {
	for _, p := range allPlugins {
		p.OnFailure(s.ToP())
	}
}

func SelectPlugin(name string) plugin.PluginActions {
	for _, p := range allPlugins {
		if p.GetInfo().Name == name {
			return p
		}
	}
	return plugin.PluginInfo{}
}

func (s *Service) PluginFailures() []*plugin.Failure {
	var failed []*plugin.Failure
	for _, f := range s.Failures {
		fail := &plugin.Failure{
			f.Id,
			f.Issue,
			f.Service,
			f.CreatedAt,
			f.Ago,
		}
		failed = append(failed, fail)
	}
	return failed
}

func (s *Service) ToP() *plugin.Service {
	out := &plugin.Service{
		s.Id,
		s.Name,
		s.Domain,
		s.Expected,
		s.ExpectedStatus,
		s.Interval,
		s.Method,
		s.Port,
		s.CreatedAt,
		s.Data,
		s.Online,
		s.Latency,
		s.Online24Hours,
		s.AvgResponse,
		s.TotalUptime,
		s.PluginFailures(),
	}
	return out
}
