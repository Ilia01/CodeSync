"use client";
import { CodeEditor } from "@/components/Editor";

export default function Home() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-r from-purple-400 to-blue-500">
      <h1 className="text-6xl font-bold text-white drop-shadow-md">
        Code Editor
      </h1>
      <CodeEditor />
    </div>
  );
}
