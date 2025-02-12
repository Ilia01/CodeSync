"use client";

import { PanelGroup, PanelResizeHandle } from "react-resizable-panels";
import { useEditor } from "@/hooks/useEditor";
import { useFiles } from "@/hooks/useFiles";
import { Editor } from "./CodeEditor";
import { Sidebar } from "./Sidebar";

export const CodeEditor = () => {
  const { currentFile, setCurrentFile, files, setFiles } = useFiles();
  const { editorRef, handleEditorMount } = useEditor(currentFile, files);

  return (
    <PanelGroup direction="horizontal">
      <Sidebar
        files={files}
        currentFile={currentFile}
        setCurrentFile={setCurrentFile}
      />
      <PanelResizeHandle className="w-2 bg-gray-700 cursor-ew-resize" />
      <Editor
        editorRef={editorRef}
        handleEditorMount={handleEditorMount}
        currentFile={currentFile}
        files={files}
      />
    </PanelGroup>
  );
};
