i=0
while [ $i -lt 115 ]; do
  jq ".[$i]" real-logs.txt >/tmp/js1.json
  # mosquitto_pub -h localhost -p 1883 -t arrebufo -q 0 -f ./td.json
  mosquitto_pub -h x.apps.p2code-uop-openshift.rh-horizon.eu -p 30883 -t EdgeXEvents -q 0 -f /tmp/js1.json
  i=$((i + 1))
  sleep 5
done
