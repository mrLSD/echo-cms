GVT=$( which gvt | grep -c ".*" )
GOVENDOR=$( which govendor | grep -c ".*" )

if [ $GVT -eq 0 ] || [ $GOVENDOR -eq 0 ]
then
	echo "gvt not found. Please install it."
	exit;
fi

if [ $GOVENDOR -eq 0 ]
then
	echo "govendor not found. Please install it."
	exit;
fi

rm -rf vendor
govendor init

for Package in $(govendor list +m | grep -o "git.*")
do
	gvt fetch $Package
done