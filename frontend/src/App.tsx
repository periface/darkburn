import { useEffect, useState } from 'react';
import './App.css';
import { StartApp, CopyToClipboard, OpenInExplorer } from "../wailsjs/go/main/App";
import { models } from '../wailsjs/go/models';
import Menu from './components/Menu';
import { ToastContainer, toast } from 'react-toastify';

function App() {
    const [result, setResult] = useState<models.Result>();
    const [files, setFiles] = useState<models.FileList[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const updateFileNames = (result: models.Result) => setResult(result);
    function start_app() {
        setLoading(true);
        StartApp().then(data => {
            updateFileNames(data);
            setFiles(data.Files);
        }).catch(err => {
            console.log(err);
        }).finally(() => {
            setLoading(false);
        });
    }
    function select_folder(folder: string) {
        console.log("Selecting folder");
        OpenInExplorer(folder).then(data => {
            console.log(data);
        });
    }

    function copy_clipboard(route: string, e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        e.preventDefault();
        const btn = e.currentTarget;
        btn.classList.add("cursor-not-allowed");
        btn.classList.add("opacity-50");
        CopyToClipboard(route).then(data => {
            toast.success("Copiado al portapapeles", {
                position: "bottom-center",
                autoClose: 2000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });

            btn.classList.remove("cursor-not-allowed");
            btn.classList.remove("opacity-50");
        }).catch(err => {
            console.log(err);
            toast.error("Error al copiar al portapapeles", {
                position: "bottom-center",
                autoClose: 2000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
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
        start_app();
    }, []);

    return (
        <div id="App">
            <div className="w-full">
                <Menu onsearch={search_files} />
            </div>
            {loading && <div className="w-full h-full flex justify-center items-center">
                <div className="w-1/2 h-1/2">
                    <h1 className="text-2xl text-center text-pink-900 font-bold">Cargando...</h1>
                </div>
            </div>
            }
            {!loading && <div className="grid grid-cols-1">
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

                                <button onClick={(e) => {
                                    copy_clipboard(file.AbsolutePath, e);
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
            }
        </div>

    )
}


export default App
