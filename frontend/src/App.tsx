import { useEffect, useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { GetFiles, CopyToClipboard, OpenInExplorer } from "../wailsjs/go/main/App";
import { main } from '../wailsjs/go/models';

function App() {
    const [result, setResult] = useState<main.Result>();
    const [path, setPath] = useState<string>("");
    const updateFileNames = (result: main.Result) => setResult(result);
    function get_files() {
        GetFiles().then(data => {
            updateFileNames(data);
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



    useEffect(() => {
        get_files();
    }, []);

    return (
        <div id="App">
            <h1 className="text-red-800 text-4xl font-bold">DarkBurn</h1>
            <div className="grid grid-cols-1">
                <div className="grid grid-cols-2">
                    {result?.Files.map((file, index) => (
                        <div key={index} className="p-2 border-2 border-solid">
                            <p><a href="#" onClick={() => {
                                console.log(file.AbsolutePath);
                            }}>{file.Name}</a>
                                <img src={file.AbsolutePath} alt="Preview" onError={
                                    e => {
                                        e.currentTarget.remove();
                                    }
                                } />
                            </p>
                            <button onClick={() => {
                                select_folder(file.AbsolutePath);
                            }} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                                Abrir carpeta
                            </button>

                            <button onClick={() => {
                                copy_clipboard(file.AbsolutePath);
                            }} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                                Copiar a clipboard
                            </button>
                        </div>
                    ))}
                </div>
            </div>

        </div>
    )
}


export default App
