"use client";
import { useState } from "react";

const initialFiles = {
  "index.ts": "// Start coding in TypeScript",
  "app.tsx": "// App Component",
  "utils.ts": "// Utility functions",
};

export const useFiles = () => {
  const [currentFile, setCurrentFile] = useState("index.ts");
  const [files, setFiles] = useState<Record<string, string>>(initialFiles);

  return { currentFile, setCurrentFile, files, setFiles };
};
