import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { useForm } from "@mantine/hooks";
import { useState } from "react";
import { KeyedMutator } from "swr";
import "./ListBox.css";

export interface List {
  id: number;
  title: string;
  body: string;
}

function ListBox() {
  return (
    <>
      <div className="center">
        <p></p>
      </div>
    </>
  );
}

export default ListBox;
