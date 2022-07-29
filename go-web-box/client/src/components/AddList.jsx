import { useState } from "react";
import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";

function AddTodo() {
    const [open, setOpen] = useState(false);
    return (
        <div>
            <Modal opened={open} onClose={() => setOpen(false)} title="Create Todo">
                <form>
                    <TextInput
                        required
                        mb={12}
                        label="Todo"
                        placeholder="What do you want to do?"
                    />
                    <Textarea
                        required
                        mb={12}
                        label="Body"
                        placeholder="Tell me more..."
                    />

                    <Button type="submit">Create todo</Button>
                </form>
            </Modal>

            <Group position="center">
                <Button fullWidth mb={12} onClick={() => setOpen(true)}>
                    ADD TODO
                </Button>
            </Group>
        </div>
    );
}

export default AddTodo;