package routes

import (
	"Backend-Golang/src/controllers/addcontrol"
	"Backend-Golang/src/controllers/bagcontrol"
	"Backend-Golang/src/controllers/catecontrol"
	"Backend-Golang/src/controllers/checkcontrol"
	"Backend-Golang/src/controllers/coscontrol"
	"Backend-Golang/src/controllers/paycontrol"
	"Backend-Golang/src/controllers/prodcontrol"
	"Backend-Golang/src/controllers/selcontrol"
	"Backend-Golang/src/controllers/usercontrol"

	// "Backend-Golang/src/controllers/usercontrol"
	"Backend-Golang/src/middleware"
	"fmt"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Router() {
	// membuat web static untuk upload image
	fileServer := http.FileServer(http.Dir("src/uploads"))

	http.Handle("/img/", http.StripPrefix("/img/", fileServer))

	helmet := helmet.Default()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello World!")
	})
	//ini testing untuk melihat,dan add ,
	http.Handle("/costumers", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(coscontrol.Costumers))))
	http.Handle("/costumer/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(coscontrol.Costumer))))
	http.Handle("/sellers", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(selcontrol.Sellers))))
	http.Handle("/seller/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(selcontrol.Seller))))
	// http.Handle("/product/", http.HandlerFunc(prodcontrol.Product))
	// http.Handle("/products", http.HandlerFunc(prodcontrol.Products))
	// http.Handle("/search/", http.HandlerFunc(prodcontrol.SearchProduct))
	// http.Handle("/orders", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ordcontrol.Orders))))
	// http.Handle("/order/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ordcontrol.Order))))
	// http.Handle("/users", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontrol.Users))))
	// http.Handle("/user/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontrol.User))))
	// http.Handle("/addresses", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(addcontrol.Addresses))))
	// http.Handle("/address/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(addcontrol.Address))))

	//Route User
	// http.Handle("/register-seller", middleware.JwtMiddleware(http.HandlerFunc(usercontrol.RegisterSeller)))
	// http.Handle("/register-costumer", middleware.JwtMiddleware(http.HandlerFunc(usercontrol.RegisterCostumer)))
	// http.Handle("/login", middleware.JwtMiddleware(http.HandlerFunc(usercontrol.Login)))
	// http.Handle("/upload", middleware.JwtMiddleware(http.HandlerFunc(prodcontrol.Handle_upload)))

	/*Xss*/

	/*JWT*/
	//Route User
	http.Handle("/register-seller", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontrol.RegisterSeller))))
	http.Handle("/register-customer", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontrol.RegisterCustomer))))
	// http.Handle("/login", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(usercontrol.Login))))
	http.Handle("/login", helmet.Secure(http.HandlerFunc(usercontrol.Login)))

	http.Handle("/update-seller/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontrol.UpdateSeller))))
	http.Handle("/update-customer/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontrol.UpdateCustomer))))

	http.Handle("/users", middleware.JwtMiddleware(http.HandlerFunc(usercontrol.Users)))
	http.Handle("/user/", middleware.JwtMiddleware(http.HandlerFunc(usercontrol.User)))

	//Route Refresh Token
	http.Handle("/refresh-token", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(usercontrol.RefreshToken))))

	//Route Product + Upload + Search
	// http.Handle("/products", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(prodcontrol.Products)))) //pagination
	// http.Handle("/products", helmet.Secure((http.HandlerFunc(prodcontrol.Products))))//pagination
	http.Handle("/products", helmet.Secure((http.HandlerFunc(prodcontrol.Products))))
	http.Handle("/selectproducts", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(prodcontrol.SelectProducts))))
	// http.Handle("/product/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(prodcontrol.Product))))
	http.Handle("/product/", helmet.Secure((http.HandlerFunc(prodcontrol.Product))))
	http.Handle("/product-search/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(prodcontrol.SearchProduct))))

	// http.Handle("/upload", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(prodcontrol.Handle_upload))))
	http.Handle("/upload", helmet.Secure((http.HandlerFunc(prodcontrol.Handle_upload))))

	//Route Bag (Rename Cart)
	http.Handle("/bags", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(bagcontrol.Bags))))
	http.Handle("/bag/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(bagcontrol.Bag))))

	//Route Address
	http.Handle("/addresses", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(addcontrol.Addresses))))
	http.Handle("/address/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(addcontrol.Address))))

	//Route Checkout
	http.Handle("/checkouts", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(checkcontrol.Checkouts))))
	http.Handle("/checkout/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(checkcontrol.Checkout))))

	//Route Payment
	http.Handle("/payments", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(paycontrol.Payments))))
	http.Handle("/payment/", helmet.Secure(middleware.JwtMiddleware(http.HandlerFunc(paycontrol.Payment))))

	// Route Category
	http.Handle("/categories", helmet.Secure((http.HandlerFunc(catecontrol.Categories))))
	http.Handle("/category/", helmet.Secure((http.HandlerFunc(catecontrol.Category))))
}
