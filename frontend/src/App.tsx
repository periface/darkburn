import { useEffect, useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { GetFiles, CopyToClipboard, OpenInExplorer } from "../wailsjs/go/main/App";
import { main } from '../wailsjs/go/models';
import Menu from './components/Menu';
function App() {
    const [result, setResult] = useState<main.Result>();
    const [path, setPath] = useState<string>("");
    const [files, setFiles] = useState<main.FileList[]>([]);
    const updateFileNames = (result: main.Result) => setResult(result);
    function get_files() {
        GetFiles().then(data => {
            updateFileNames(data);
            setFiles(data.Files);
        });
    }
    function select_folder(folder: string) {
        console.log("Selecting folder");
        OpenInExplorer(folder).then(data => {
            console.log(data);
        });
    }

    function copy_clipboard(route: string) {
        console.log("Copying to clipboard");
        CopyToClipboard(route).then(data => {
            console.log(data);
        })
    }

    function search_files(search: string) {
        console.log("Searching files");
        if (result?.Files?.length) {
            let filtered_files = result?.Files.filter((file) => {
                return file.Name.toUpperCase().includes(search.toUpperCase()) ||
                    file.Extension.toUpperCase().includes(search.toUpperCase())
                    || file.AbsolutePath.toUpperCase().includes(search.toUpperCase());
            });
            setFiles(filtered_files);
        }
    }

    useEffect(() => {
        get_files();
    }, []);

    return (
        <div id="App">
            <div className="w-full">
                <Menu onsearch={search_files} />
            </div>
            <div className="grid grid-cols-1">
                <div className="grid grid-cols-3">
                    {files.map((file, index) => (
                        <div key={index} className="p-2 relative group">

                            <div className='z-50 opacity-0 group-hover:opacity-100 absolute top-1/2 right-1/2 bg-opacity-50
                                duration-300 ease-in-out transition-all delay-200
                                transform translate-x-1/2 -translate-y-1/2 w-3/4'>
                                <h4 className="font-bold text-lg text-center text-blue-900">{file.Name}</h4>
                                <button onClick={() => {
                                    select_folder(file.AbsolutePath);
                                }} className="bg-pink-700 hover:bg-pink-900 rounded-full text-white font-bold py-2 px-4 w-full text-sm mb-2">
                                    Abrir
                                </button>

                                <button onClick={() => {
                                    copy_clipboard(file.AbsolutePath);
                                }} className="bg-blue-500 hover:bg-blue-700 rounded-full text-white font-bold py-2 px-4 w-full text-sm">
                                    Copiar
                                </button>
                            </div>
                            <div className='border-sky-500 shadow-lime-50 grid grid-cols-1 align-middle items-center justify-items-center'>
                                <div className=''>
                                    <p className="m-0"><a href="#" className='cursor-pointer' onClick={() => {
                                        console.log(file.AbsolutePath);
                                    }}>{file.Name}</a>
                                        <img src={file.AbsolutePath} alt="Preview"
                                            className="opacity-100 group-hover:opacity-20
                                            group-hover:w-2/4 m-auto group-hover:m-auto object-contain
                                            object-center overflow-hidden w-64 h-64
                                            duration-500 ease-in-out transition-all"
                                            onError={
                                                e => {
                                                    e.currentTarget.remove();
                                                }
                                            } />

                                    </p>
                                </div>
                                <div className='w-1/2 m-auto'>
                                    <small className='text-gray-500 text-xs'>{file.Name}</small>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>

        </div >
    )
}


export default App
