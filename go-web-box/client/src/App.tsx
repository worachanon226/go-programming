import { useState } from "react";
import AddList, { List } from "./components/AddList";
import ListBox from "./components/ListBox";
import "./App.css";
import { mutate } from "swr";

function App() {
  return (
    <div>
      <AddList mutate={mutate}></AddList>
      <ListBox></ListBox>
    </div>
  );
}

export default App;
