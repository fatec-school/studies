#!/bin/bash

RESOURCE_GROUP="MyResourceGroup"
LOCATION="eastus"
PASSWORD="YourPassword123!"

make_vm() {
    local NOME_VM=$1
    echo "Criando VM: $NOME_VM"

    az vm create \
        --resource-group $RESOURCE_GROUP \
        --name $NOME_VM \
        --image UbuntuLTS \
        --admin-username azureuser \
        --admin-password $PASSWORD \
        --location $LOCATION \
        --output none \
        --no-wait       
}

VMS=(vm1 vm2)

for nome in "${VMS[@]}"; do
    make_vm $nome
done

echo "Todas as VMs estão sendo criadas em segundo plano."