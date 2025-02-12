import { Panel } from "react-resizable-panels";
import MonacoEditor from "@monaco-editor/react";
import type * as monaco from "monaco-editor";
interface EditorProps {
  editorRef: React.MutableRefObject<monaco.editor.IStandaloneCodeEditor | null>;
  handleEditorMount: (editor: monaco.editor.IStandaloneCodeEditor) => void;
  currentFile: string;
  files: Record<string, string>;
}

export const Editor = ({
  editorRef,
  handleEditorMount,
  currentFile,
  files,
}: EditorProps) => (
  <Panel defaultSize={80} className="bg-gray-800">
    <MonacoEditor
      height="100vh"
      theme="vs-dark"
      defaultLanguage="typescript"
      defaultValue={files[currentFile]}
      onMount={handleEditorMount}
      options={{
        automaticLayout: true,
        minimap: { enabled: true },
        fontSize: 16,
        lineNumbers: "on",
        scrollbar: { vertical: "auto", horizontal: "auto" },
        wordWrap: "on",
      }}
    />
  </Panel>
);
