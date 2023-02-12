const os = require('os');
const path = require('path');
const process = require("child_process");
const fs = require("fs");
const config = require("../config/build.json");

const GRPC_PROJECT = config.grpcPackage;
const GO_PROJECT = config.goPackage;



const isUndefinedEnv = (env) => {
	return typeof env === 'undefined';
};


const makeCommand = (grpcRoot, goPath, fileSource) => {
	return `protoc -I ${grpcRoot} --go_out=${goPath} --go-grpc_out=${goPath}  ${fileSource}`;
};

const getFileList = (dir) => {
	return fs.readdirSync(dir).filter((value) => {
		return path.extname(value) == ".proto";	
	});
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

	grpcP = path.resolve(grpcP);
	goP = path.resolve(goP);

	let fileSource = getFileList(grpcP).reduce((acc,current)=>{return `${acc} ${current}`},"");
	
	let cmd = makeCommand(grpcP, goP, fileSource);
	
	try {
		process.execSync(cmd);
	}catch(e){
		console.log(`${e.stderr}`);
		console.log("fail protoc grpc compile fail");
		return;
	}
	
	console.log("protoc grpc compile done!\n");

};

buildMain(GRPC_PROJECT,GO_PROJECT);
