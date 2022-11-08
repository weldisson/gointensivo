package entity // nome do pacote é o nome da pasta

import "errors"

type Order struct { // um tipo Order que é um struct semelhante a uma classe
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// *Order e &Order é um ponteiro, irá ser falado mais na proxima aula.

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	// o go garante que tipagem
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()
	if err != nil {
		return nil, err
	}

	return order, nil
}

// uma validação atrelada ao struct Order que retorna um error
func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
