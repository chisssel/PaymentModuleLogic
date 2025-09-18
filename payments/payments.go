package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}
type PaymentModule struct {
	paymentMethod PaymentMethod
	paymentsInfo  map[int]PaymentInfo
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
		paymentsInfo:  make(map[int]PaymentInfo),
	}
}

// Метод оплаты
// Принимает:
// 1. Описание проводимой оплаты
// 2. Сумму оплаты
// Возвращает:
// 1. ID проведенной операции
func (p *PaymentModule) Pay(description string, usd int) int {
	// 1. Проводить оплату
	// 2. Получать id проведенной оплаты
	id := p.paymentMethod.Pay(usd)

	// 3. Сохранять информацию о проведенной операции
	// - описание операции
	// - сколько было потрачено
	// - отменённая ли операция
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	p.paymentsInfo[id] = info

	// 4. Возвращать id проведенной операции
	return id
}

// Метод отмены
// Принимает:
// 1. ID операции
// Возвращает:
// - Ничего не возвращает
func (p *PaymentModule) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.paymentsInfo[id] = info
}

// Метод инфо одной операции
// Принимает:
// 1. ID операции
// Возвращает:
// 1. Информацию о проведенной операции
func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// Метод инфо всех операций
// Принимает:
// - Ничего не принимает
// Возвращает:
// - Информацию о всех проведенных операциях
func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))

	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}
	return tempMap
}
