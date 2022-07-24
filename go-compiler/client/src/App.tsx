import useSWR from "swr";
import * as React from "react";
import "./App.css";
import LoadingSpinner from "./loading";

export const ENDPOINT = "http://localhost:8080";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.text);

function App() {
  const [langugue, setLangugue] = React.useState("cpp");
  const [code, setCode] = React.useState("");
  const [isLoading, setIsLoading] = React.useState(false);
  const [output, setOutput] = React.useState("");

  const { data, mutate } = useSWR(`servercheck`, fetcher);

  async function complie() {
    setIsLoading(true);
    const input = await fetch(`${ENDPOINT}/run/${langugue}`, {
      method: "POST",
      headers: {
        "Content-Type": "text/plain; charset=utf-8",
      },
      body: code,
    });

    console.log("Input success");

    const output = await fetch(`${ENDPOINT}/output/${langugue}`, {
      method: "GET",
      headers: {
        "Content-Type": "text/plain; charset=utf-8",
      },
    });
    console.log("Output success");
    if (output.ok) {
      let s = await output.text();
      setOutput(s);
      console.log(s);
    } else {
      let e = output.status;
      console.log(e);
    }
    setIsLoading(false);
  }

  const TextArea = (
    <textarea
      rows={20}
      cols={75}
      value={code}
      onChange={(e) => {
        setCode(e.target.value);
      }}
    ></textarea>
  );

  return (
    <div className="App">
      <h1>Online Code Complier</h1>
      <div>
        <label>Langugue: </label>
        <select
          value={langugue}
          onChange={(e) => {
            setLangugue(e.target.value);
          }}
        >
          <option value="cpp">C++</option>
          <option value="py">Python</option>
        </select>
      </div>
      <br />
      {isLoading ? <LoadingSpinner /> : TextArea}
      <br />
      <button onClick={complie} disabled={isLoading}>
        Submit
      </button>
      <p>{output}</p>
    </div>
  );
}

export default App;
