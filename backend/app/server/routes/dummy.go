package routes

import (
	"fmt"
	"math/rand"

	"github.com/valyala/fasthttp"
)

var msg []string = []string{
	"What does an upset chef make food with?\nAngrydients.",
	"What is the sushi chef's dream car?\nRolls Rice",
	"I was eavesdropping on two indian chefs talking...\nIt was a dhal conversation",
	"What did the pastry chef say to his unsupportive father?\nDoughnut hole me back.",
	"Fdhvdu vdodg",
}

// Dummy route that just shows some random messages and jokes
func Dummy(ctx *fasthttp.RequestCtx, _ *fasthttp.Client, _ int) {
	ctx.SetContentType("text/plain")
	ctx.SetStatusCode(fasthttp.StatusTeapot)
	fmt.Fprintf(ctx, msg[rand.Intn(len(msg))])
}
