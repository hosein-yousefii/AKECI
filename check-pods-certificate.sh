

APPS=($(kubectl get pods --show-labels --no-headers|awk '{print $6}'|cut -d= -f2|cut -d, -f1|xargs))

for i in ${APPS[@]}
do
	
	RENEW_DATE=$(date +%s -d "$(kubectl describe certificate $i|grep "Renewal Time"|awk '{print $3}')")
	DATE_NOW=$(date +%s)


	if [[ $DATE_NOW -lt $RENEW_DATE ]]
	then

		kubectl delete certificate $i
		kubectl delete secret $i

	fi
done
