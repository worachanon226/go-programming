import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { useForm } from "@mantine/hooks";
import { useState } from "react";
import { KeyedMutator } from "swr";
import "./AddList.css";

export interface List {
  id: number;
  title: string;
  body: string;
}

function AddList({ mutate }: { mutate: KeyedMutator<List[]> }) {
  const [open, setOpen] = useState(false);

  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

  return (
    <>
      <Group>
        <Button onClick={() => setOpen(true)}>ADD LIST</Button>
      </Group>

      <Modal opened={open} onClose={() => setOpen(false)} title="ADD LIST">
        <form>
          <TextInput
            required
            mb={12}
            label="Title"
            {...form.getInputProps("title")}
          ></TextInput>
          <Textarea
            required
            mb={12}
            label="Body"
            {...form.getInputProps("body")}
          ></Textarea>

          <Button type="submit">Create list</Button>
        </form>
      </Modal>
    </>
  );
}

export default AddList;
