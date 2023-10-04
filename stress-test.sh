GATLING_BIN_DIR=../gatling/bin
GATLING_WORKSPACE=../gatling
RESULTS_WORKSPACE=../gatling/user-files/results

competidores=("DeveloperArthur")

for competidor in ${competidores[@]}; do
(
    diretorio="app"
    echo "======================"
    echo $competidor

    cd $diretorio
    echo "iniciando e logando execução da API"
    mkdir "$RESULTS_WORKSPACE/$competidor"
    sudo docker-compose up -d --build
    sudo docker-compose logs > "$RESULTS_WORKSPACE/$competidor/docker-compose.logs"
    echo "pausa de 15 segundos para startup pra API"
    sleep 15
    echo "iniciando teste"
    sudo sh $GATLING_BIN_DIR/gatling.sh -rm local -s RinhaBackendSimulation \
        -rd $competidor \
        -rf "$RESULTS_WORKSPACE/$competidor" \
        -sf $GATLING_WORKSPACE/user-files/simulations \
        -rsf $GATLING_WORKSPACE/user-files/resources
    echo "teste finalizado"
    echo "fazendo request e salvando a contagem de pessoas"
    curl -H -i -N "http://localhost:9999/contagem-pessoas"
    sudo curl -v "http://localhost:9999/contagem-pessoas" > "$RESULTS_WORKSPACE/$competidor/contagem-pessoas.log"
    echo "cleaning up do docker"
    sudo docker-compose rm -f
    sudo docker-compose down
)
done
