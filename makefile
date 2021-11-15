publish:
	git add . 
	git commit -m $m
	git tag $v
	git push origin $v
	GOPROXY=proxy.golang.org go list -m github.com/bankonly/gomodel@$v