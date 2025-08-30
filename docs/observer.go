type Observer interface {
	Update(data string)
}

type Subject struct {
 observers []Observer
}

func (s *Subject) Register(o Observer) {
 s.observers = append(s.observers, o)
}

func (s *Subject) Notify(data string) {
 for _, o := range s.observers {
  o.Update(data)
 }
}

type Logger struct{}
func (l *Logger) Update(data string) {
 fmt.Println("Logger received:", data)
}
subject := &Subject{}
logger := &Logger{}
subject.Register(logger)
subject.Notify("New event occurred")