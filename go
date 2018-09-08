snippet v
	${1} := ${0}

snippet vr
	var ${1} = ${0}

snippet vrs
	var (
		${1} = ${0}
	)

snippet vi
	var ${1} int = ${0}

snippet vs
	var ${1} string = ${0}

snippet vfl
	var ${1} float32 = ${0}

snippet vflb
	var ${1} float64 = ${0}

snippet vbl
	var ${1} bool = ${0}

snippet vbt
	var ${1} byte = ${0}

snippet cn
	const ${1} = ${0}

snippet cni
	const ${1} int = ${0}

snippet cns
	const ${1} string = ${0}

snippet cnfl
	const ${1} float32 = ${0}

snippet cnflb
	const ${1} float64 = ${0}

snippet cnbl
	const ${1} bool = ${0}

snippet cnbt
	const ${1} byte = ${0}

snippet i
	int 

snippet s
	string 

snippet fl
	float

snippet flb
	float32

snippet bl
	bool

snippet bt
	byte

snippet pa
	package ${1:main}
	
snippet fn
	func ${1}(${2}) {
		${0}
	}

snippet fnm
	func main() {
		${0}
	}

snippet fne
	func ${1}(${2}) ${3} {
		${0}
	}

snippet rt
	return  

snippet rs
	result

snippet pf
	fmt.Printf(${0})

snippet pfs
	fmt.Printf("${0}")

snippet pfv
	fmt.Printf("%${1}", ${0})
	

snippet pl
	fmt.Println("${0}")

snippet lf
	log.Printf("%${1}", ${0})

snippet ll
	log.Println("${0}")

snippet if
	if ${1} {
		${0}
	}

snippet ife
	if ${1} == ${2} {
		${0}
	}

snippet ifne
	if ${1} != ${2} {
		${.3}
	}

snippet ifb
	if ${1} > ${2} {
		${0}
	}

snippet ifbe
	if ${1} >= ${2} {
		${0}
	}

snippet ifs
	if ${1} < ${2} {
		${0}
	}

snippet ifse
	if ${1} <= ${2} {
		${0}
	}

snippet elif
	else if ${1} {
		${0}
	}

snippet elife
	else if ${1} == ${2} {
		${0}
	}

snippet elifne
	else if ${1} != ${2} {
		${0}
	}

snippet elifb
	else if ${1} > ${2} {
		${0}
	}

snippet elifbe
	else if ${1} >= ${2} {
		${0}
	}

snippet elifs
	else if ${1} < ${2} {
		${0}
	}

snippet elifse
	else if ${1} <= ${2} {
		${0}
	}

snippet el
	else {
		${0}
	}

snippet f
	false 

snippet t
	true 

snippet for
	for ${1} {
		${0}
	}

snippet fori
	for ${1:i} := ${2:0}; $1 ${3:<} ${4}; $1${5:++} {
		${0}
	}
