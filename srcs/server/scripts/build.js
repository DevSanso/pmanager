const os = require('os');
const path = require('path');
const process = require("child_process");

const GRPC_PROJECT = process.env.PMANGER_GRPC_PROJECT;
const GO_PROJECT = process.env.PMANAGER_GO_PROJECT;



const isUndefinedEnv = (env) => {
	return typeof env === 'undefined';
};


const makeCommand = (grpcPath,goPath, entry) => {
	return `protoc -I ${grpcPath} --go_out=${goPath} ${entry}`;
};


const buildMain = (grpc_p_env,go_p_env) => {
	let grpcP = grpc_p_env;
	let goP = go_p_env;

	if (isUndefinedEnv(grpcP)) {
		throw "grpc project env (PMANAGER_GRPC_PROJECT) not define";
	}
	if (isUndefinedEnv(goP)) {
		throw "go project (PMANAGER_GO_PROJECT) not define";
	}
	
	const entryFiles = "private.proto  public.proto";
	const cmd = makeCommand(grpcP,goP,entryFiles);
	
	const result = execSync(cmd);
	
	console.log("protoc compile done!\n");

};


buildMain(GRPC_PROJECT,GO_PROJECT);
