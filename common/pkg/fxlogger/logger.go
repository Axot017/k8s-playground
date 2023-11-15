package fxlogger

import (
	"log/slog"
	"strings"

	"go.uber.org/fx/fxevent"
)

type SlogLogger struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *SlogLogger {
	return &SlogLogger{logger: logger}
}

func (l *SlogLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logger.Info("OnStart hook executing",
			slog.String("callee", e.FunctionName),
			slog.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.logger.Error("OnStart hook failed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("error", e.Err.Error()),
			)
		} else {
			l.logger.Info("OnStart hook executed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.logger.Info("OnStop hook executing",
			slog.String("callee", e.FunctionName),
			slog.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.logger.Error("OnStop hook failed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("error", e.Err.Error()),
			)
		} else {
			l.logger.Info("OnStop hook executed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.logger.Error("error encountered while applying options",
				slog.String("type", e.TypeName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
				slog.String("error", e.Err.Error()))
		} else {
			l.logger.Info("supplied",
				slog.String("type", e.TypeName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
			)
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.logger.Info("provided",
				slog.String("constructor", e.ConstructorName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
				slog.String("type", rtype),
				slog.Bool("private", e.Private),
			)
		}
		if e.Err != nil {
			l.logger.Error("error encountered while applying options",
				slog.String("module", e.ModuleName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("error", e.Err.Error()))
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.logger.Info("replaced",
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
				slog.String("type", rtype),
			)
		}
		if e.Err != nil {
			l.logger.Error("error encountered while replacing",
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
				slog.String("error", e.Err.Error()))
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.logger.Info("decorated",
				slog.String("decorator", e.DecoratorName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
				slog.String("type", rtype),
			)
		}
		if e.Err != nil {
			l.logger.Error("error encountered while applying options",
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName),
				slog.String("error", e.Err.Error()))
		}
	case *fxevent.Run:
		if e.Err != nil {
			l.logger.Error("error returned",
				slog.String("name", e.Name),
				slog.String("kind", e.Kind),
				slog.String("module", e.ModuleName),
				slog.String("error", e.Err.Error()),
			)
		} else {
			l.logger.Info("run",
				slog.String("name", e.Name),
				slog.String("kind", e.Kind),
				slog.String("module", e.ModuleName),
			)
		}
	case *fxevent.Invoking:
		l.logger.Info("invoking",
			slog.String("function", e.FunctionName),
			slog.String("module", e.ModuleName),
		)
	case *fxevent.Invoked:
		if e.Err != nil {
			l.logger.Error("invoke failed",
				slog.String("error", e.Err.Error()),
				slog.String("stack", e.Trace),
				slog.String("function", e.FunctionName),
				slog.String("module", e.ModuleName),
			)
		}
	case *fxevent.Stopping:
		l.logger.Info("received signal",
			slog.String("signal", strings.ToUpper(e.Signal.String())))
	case *fxevent.Stopped:
		if e.Err != nil {
			l.logger.Error("stop failed", slog.String("error", e.Err.Error()))
		}
	case *fxevent.RollingBack:
		l.logger.Error("start failed, rolling back", slog.String("error", e.StartErr.Error()))
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.logger.Error("rollback failed", slog.String("error", e.Err.Error()))
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.logger.Error("start failed", slog.String("error", e.Err.Error()))
		} else {
			l.logger.Info("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.logger.Error("custom logger initialization failed", slog.String("error", e.Err.Error()))
		} else {
			l.logger.Info("initialized custom fxevent.Logger", slog.String("function", e.ConstructorName))
		}
	default:
		l.logger.Warn("unknown event", slog.Any("event", event))
	}
}
