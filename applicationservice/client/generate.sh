parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

# Generate User Proto
echo -e "> \e[33mGenerating \e[39muser proto.."
protoc  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative userclient/generatedclient/*.proto
echo -e "> \e[92mGenerated \e[39muser proto."

# Generate Cache Proto
echo -e "> \e[33mGenerating \e[39mcache proto.."
protoc  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative cacheclient/generatedclient/*.proto
echo -e "> \e[92mGenerated \e[39mcache proto."


printf "\n"
read -n 1 -s -r -p "Press any key to continue"