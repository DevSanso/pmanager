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
	return `protoc -I ${grpcRoot} --go_out=${goPath} ${fileSource}`;
};

const getFileList = (dir) => {
	return fs.readdirSync(dir).filter((value) => {
		return path.extname(value) == ".proto";	
	});
};

const getProtoPackageDir = (grpcPath) => {
	const reqInclude = path.join(grpcPath,"message/request");
	const resInclude = path.join(grpcPath,"message/response");
	const commonInclude = path.join(grpcPath,"message/common");

	return {
		req : reqInclude,
		res : resInclude,
		common : commonInclude,
		root : grpcPath
	};
}

function buildOneDir(element) {
	let packageDir = this.packageDir;
	let dir = Object.getOwnPropertyDescriptor(packageDir,element).value;
	
	let fileSource = getFileList(dir)
		.map(function(value){return path.join(String(this),value);},dir)
		.reduce(function(acc,current){return `${acc} ${current}`;},"");

	let cmd = makeCommand(packageDir.root, this.goP, fileSource);
	
	try {
		process.execSync(cmd);
	}catch(e){
		console.log(`${e.stderr}`);
		console.log("fail protoc compile fail");
		return;
	}
	console.log(`protoc compile package ${element} !\n`);
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
	
	const packageDir = getProtoPackageDir(grpcP);
	
	const forEachThis = {
		goP : goP,
		packageDir : packageDir
	};
	
	
	Object.keys(packageDir).forEach(buildOneDir, forEachThis);
	console.log("protoc compile done!\n");

};


buildMain(GRPC_PROJECT,GO_PROJECT);
