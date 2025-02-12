import { Panel } from "react-resizable-panels";

interface SidebarProps {
  files: Record<string, string>;
  currentFile: string;
  setCurrentFile: (file: string) => void;
}

export const Sidebar = ({
  files,
  currentFile,
  setCurrentFile,
}: SidebarProps) => (
  <Panel defaultSize={20} minSize={15} maxSize={30} className="bg-gray-900 p-4">
    <h3 className="text-white font-bold"> Files</h3>
    {Object.keys(files).map((file) => (
      <div
        key={file}
        className={`cursor-pointer p-2 text-gray-300 ${
          currentFile === file ? "bg-gray-700" : ""
        }`}
        onClick={() => setCurrentFile(file)}
      >
        {file}
      </div>
    ))}
  </Panel>
);
