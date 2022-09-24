package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Can't find product")
	ErrCantDecodeProducts = errors.New("Can't find product")
	ErrUserIdIsNotValid   = errors.New("This user is not valid")
	ErrCantUpdateUser     = errors.New("Can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Can not remove this item from cart")
	ErrCantGetItem        = errors.New("Unable to get item from the cart")
	ErrCantBuyCartItem    = errors.New("Can not update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
