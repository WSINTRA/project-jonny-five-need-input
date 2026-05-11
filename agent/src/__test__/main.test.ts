// Want to get started with the agent side of things, particularly understanding the plugin concept. 
// What better way than an empty test file. Lets go modular. First thing I need to be able to send prompts and parse streams from llama.cpp 
import {describe, expect, it, vi} from "vitest";
import {promptCli} from '../main';
describe("promptCli", () => {
  it("takes a prompt as input and forwards it to promptParse", async () => {
    const promptParse = vi.fn();
    expect(promptCli("What times dinner?", promptParse))
    expect(promptParse).toHaveBeenCalled();

  })
})
