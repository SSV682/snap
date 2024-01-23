package logging

import "github.com/sirupsen/logrus"

const alertNameField = "alert_name"

type Entry interface {
	WithFields(fields Fields) Entry

	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Alert(alertName string, args ...interface{})
	Alertf(alertName string, format string, args ...interface{})

	AlertFatal(alertName string, args ...interface{})
	AlertFatalf(alertName string, format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}

type entryImpl struct {
	entry *logrus.Entry
}

func (i *entryImpl) WithFields(fields Fields) Entry {
	return NewEntry(i.entry.WithFields(logrus.Fields(fields)))
}

func (i *entryImpl) Trace(args ...interface{}) {
	i.entry.Trace(args...)
}

func (i *entryImpl) Tracef(format string, args ...interface{}) {
	i.entry.Tracef(format, args...)
}

func (i *entryImpl) Debug(args ...interface{}) {
	i.entry.Debug(args...)
}

func (i *entryImpl) Debugf(format string, args ...interface{}) {
	i.entry.Debugf(format, args...)
}

func (i *entryImpl) Info(args ...interface{}) {
	i.entry.Info(args...)
}

func (i *entryImpl) Infof(format string, args ...interface{}) {
	i.entry.Infof(format, args...)
}

func (i *entryImpl) Warn(args ...interface{}) {
	i.entry.Warn(args...)
}

func (i *entryImpl) Warnf(format string, args ...interface{}) {
	i.entry.Warnf(format, args...)
}

func (i *entryImpl) Error(args ...interface{}) {
	i.entry.Error(args...)
}

func (i *entryImpl) Errorf(format string, args ...interface{}) {
	i.entry.Errorf(format, args...)
}

func (i *entryImpl) Alert(alertName string, args ...interface{}) {
	i.entry.WithFields(logrus.Fields{alertNameField: alertName}).Error(args...)
}

func (i *entryImpl) Alertf(alertName string, format string, args ...interface{}) {
	i.entry.WithFields(logrus.Fields{alertNameField: alertName}).Errorf(format, args...)
}

func (i *entryImpl) AlertFatal(alertName string, args ...interface{}) {
	i.Alert(alertName, args...)
	i.entry.Logger.Exit(1)
}

func (i *entryImpl) AlertFatalf(alertName string, format string, args ...interface{}) {
	i.Alertf(alertName, format, args...)
	i.entry.Logger.Exit(1)
}

func (i *entryImpl) Fatal(args ...interface{}) {
	i.entry.Fatal(args...)
}

func (i *entryImpl) Fatalf(format string, args ...interface{}) {
	i.entry.Fatalf(format, args...)
}

func (i *entryImpl) Panic(args ...interface{}) {
	i.entry.Panic(args...)
}

func (i *entryImpl) Panicf(format string, args ...interface{}) {
	i.entry.Panicf(format, args...)
}

func NewEntry(entry *logrus.Entry) Entry {
	return &entryImpl{entry}
}
