for i in `seq 1000`
do
curl http://service1:5678/ip; 
echo '';
sleep 1; 
done;