#!/usr/bin/env node

const schema = require('../schema.json');
const fs = require('fs');
const { promisify } = require('util');
const writeFile = promisify(fs.writeFile);
const {listReleases} = require("@etclabscore/dl-github-releases");

const setRPCFlowVersionEnum = async (s) => {
  s.properties.rpcflow.enum = await listReleases("etclabscore", "rpcflow-spec");
  return s;
};

const build = async () => {
  const withVersionEnum = await setRPCFlowVersionEnum(schema);
  await writeFile("schema.json", JSON.stringify(withVersionEnum, undefined, "  "));
  console.log("wrote schema.json");
};

module.exports = {setRPCFlowVersionEnum};
if (require.main === module) {
  build();
}
