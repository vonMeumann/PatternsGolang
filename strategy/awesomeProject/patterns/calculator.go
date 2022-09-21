package patterns

type Calculator struct {
	Strategy // определение интерфейса
}

func (calc *Calculator) SetStrategy(strategy Strategy) {
	calc.Strategy = strategy // присваивание стратегии, чтобы можно было менять реализацию
}
