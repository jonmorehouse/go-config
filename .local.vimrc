" 
fu! GoTestPackage()

	let command = "go test -v ."

	call CleanShell(command)

endfunction

fu! GoCheckTestPackageVerbose() 

	let command = "go test -gocheck.v ."

	call CleanShell(command)

endfunction

fu! GoTestFile()
	
	if @% !~ ".go"

		return

	elseif @% =~ "_test.go"

		let file = @%

	else
		let file = split(@%, ".go")[0] . "_test.go"

	endif

	" now run go tests
	let command = "go test -gocheck.v " . file . " test_bootstrap.go"

	" now run this command
	call CleanShell(command)

endfunction

map <Leader>rr :call GoTestPackage()<CR>


