import * as fs from "fs";
import Dereferencer from "@json-schema-tools/dereferencer";

describe("meta-schema", () => {
  const s = fs.readFileSync("./schema.json", "utf8");
  let schema = false;
  try {
    schema = JSON.parse(s);
  } catch (e) {
    console.error(`Cannot parse meta-schema. Recieved: ${s}`);
  }

  it("is valid json", () => {
    expect(schema).toBeTruthy();
  });

  it("can be run through the dereffer", async () => {
    expect.assertions(1);
    const deref = new Dereferencer(schema);
    const dereffed = await deref.resolve();
    expect(dereffed).toBeTruthy();
  });
});
