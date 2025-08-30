type Service interface {
	Execute() string
}

type BaseService struct{}

func (b *BaseService) Execute() string {
 return "Executing base service"
}

type LoggingDecorator struct {
 Wrapped Service
}

func (l *LoggingDecorator) Execute() string {
 log.Println("Before execution")
 res := l.Wrapped.Execute()
 log.Println("After execution")
 return res
}
svc := &LoggingDecorator{Wrapped: &BaseService{}}
fmt.Println(svc.Execute())