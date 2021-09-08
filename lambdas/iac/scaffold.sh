echo "Converting this lambda folder to use Terraform."

if [ $# -eq 0 ]; then
  echo "No arguments provided, you MUST provide the app name!"
  return
fi

APP=$1

cp -f ../iac/templates/Makefile Makefile
sed -i '' "s/APP/${APP}/g" Makefile

SetUpStage() {
  cp -f ../../iac/templates/terraform.tfvars terraform.tfvars
  sed -i '' "s/STAGE/${STAGE}/g" terraform.tfvars

  cp -f ../../iac/templates/terraformer.tf terraformer.tf
  sed -i '' "s/APP/${APP}/g" terraformer.tf
  sed -i '' "s/STAGE/${STAGE}/g" terraformer.tf

  ln -s ../main.tf main.tf
  ln -s ../variables.tf variables.tf
}

mkdir -p "iac/base"
mkdir -p "iac/dev"
mkdir -p "iac/qa"
mkdir -p "iac/uat"
mkdir -p "iac/prod"
mkdir -p "handler"

cd iac
cp -f ../../iac/templates/main.tf main.tf
cp -f ../../iac/templates/variables.tf variables.tf
sed -i '' "s/APP/${APP}/g" variables.tf

cd dev
STAGE='dev'
SetUpStage
cd ../qa
STAGE='qa'
SetUpStage
cd ../uat
STAGE='uat'
SetUpStage
cd ../prod
STAGE='prod'
SetUpStage

cd ../..
