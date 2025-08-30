
var instance *Config
var once sync.Once

type Config struct {
 DatabaseURL string
}
func GetConfigInstance() *Config {
 once.Do(func() {
  instance = &Config{DatabaseURL: "postgres://localhost"}
 })
 return instance
}

func BenchmarkSingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetConfigInstance()
	}
}