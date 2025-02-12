import { useRef, useEffect } from "react";
import { WebsocketProvider } from "y-websocket";
import { Doc } from "yjs";
import { MonacoBinding } from "y-monaco";
import type * as monaco from "monaco-editor";

export const useEditor = (
  currentFile: string,
  files: Record<string, string>
) => {
  const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);

  useEffect(() => {
    if (!editorRef.current) return;

    const ydoc = new Doc();
    const provider = new WebsocketProvider(
      "ws://localhost:8080/ws/",
      "my-room",
      ydoc
    );

    const model = editorRef.current?.getModel();
    if (model) {
      const ytext = ydoc.getText(currentFile);
      new MonacoBinding(
        ytext,
        model,
        new Set([editorRef.current!]),
        provider.awareness
      );
    }

    provider.on("status", (event: { status: string }) => {
      console.log("WebSocket status:", event.status);
    });

    return () => {
      provider.destroy();
      ydoc.destroy();
    };
  }, [currentFile]);

  const handleEditorMount = (editor: monaco.editor.IStandaloneCodeEditor) => {
    editorRef.current = editor;
  };

  return { editorRef, handleEditorMount };
};
