# ponzu-go-example

Super fast JSON API admin framework.

## Documentation
For more detailed documentation, check out the [docs](https://docs.ponzu-cms.org)


## Requirements
Go 1.8+

Since HTTP/2 Server Push is used, Go 1.8+ is required. However, it is not 
required of clients connecting to a Ponzu server to make HTTP/2 requests. 

## Generate API content example

```bash
$ ponzu gen content product productId:"int" name:"string" price:"string" description:"string":richtext category:@category,name
$ ponzu gen content order orderId:"int" customer_name:"string" products:[]@product,name, price
```

## Serve

```bash
$ ponzu build
$ ponzu run
```
