import { useEffect, useState } from 'react';
import './App.css';
import { StartApp, CopyToClipboard, OpenInExplorer, GetFiles } from "../wailsjs/go/main/App";
import { models } from '../wailsjs/go/models';
import Menu from './components/Menu';
import { ToastContainer, toast } from 'react-toastify';

function App() {
    const [files, setFiles] = useState<models.FileList[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [loadingApp, setLoadingApp] = useState<boolean>(true);
    const [unfiltered_files, setUnfilteredFiles] = useState<models.FileList[]>([]);
    const [search, setSearch] = useState<string>("");
    const [filetype, setFileType] = useState<string>("todos");
    function select_folder(folder: string) {
        OpenInExplorer(folder).then(data => {
            console.log(data);
        });
    }

    function copy_clipboard(file_item: models.FileList, e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        e.preventDefault();
        const route = file_item.AbsolutePath;
        const btn = e.currentTarget;
        if (file_item.Extension !== ".svg") {
            toast.warn("DXF no soportado para copia, abrir carpeta", {
                position: "bottom-center",
                autoClose: 2000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
            });
            return;
        }
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

    useEffect(() => {
        async function search_files() {
            console.log("Searching files");
            if (loadingApp) {
                return;
            }
            console.log("Searching files", search, filetype);
            const files = await GetFiles(search, filetype)
            setFiles(files);
            setUnfilteredFiles(files);
        }
        search_files()
    }, [search, filetype, loadingApp]);
    useEffect(() => {

        async function start_app() {
            try {
                setLoading(true);
                await StartApp();
            }
            catch (err) {
                console.log(err);
            }
            finally {
                setLoadingApp(false);
                setLoading(false);
            }

        }
        start_app();

    }, []);

    return (
        <div id="App" className='max-w-full bg-black'>
            <div className="w-full">
                <Menu onsearch={txt => {
                    setSearch(txt);
                }} onselect={txt => {
                    setFileType(txt);
                }} />
            </div>
            {loading && <div className="w-full h-full flex justify-center items-center">
                <div className="w-1/2 h-1/2">
                    <h1 className="text-2xl text-center text-white-900 font-bold">Cargando...</h1>
                </div>
            </div>
            }
            {!loading && files?.length === 0 && <div className="w-full h-full flex justify-center items-center">
                <div className="w-1/2 h-1/2">
                    <h1 className="text-2xl text-center text-white-900 font-bold">No hay archivos</h1>
                </div>
            </div>
            }
            {!loading && files?.length && <div className="grid grid-cols-1 max-w-full">
                <div className="grid grid-cols-3">
                    {files.map((file, index) => (

                        <div key={index} className="p-2 relative group w-full bg-black">
                            <div className='z-50 opacity-0 group-hover:opacity-100 absolute top-1/2 right-1/2 bg-opacity-50
                                duration-300 ease-in-out transition-all delay-200
                                transform translate-x-1/2 -translate-y-1/2 w-3/4'>
                                <h4 className="font-bold text-lg text-center text-blue-900 w-full max-w-full
                                    break-words">{file.Name}</h4>
                                <button onClick={() => {
                                    select_folder(file.AbsolutePath);
                                }} className="bg-pink-700 hover:bg-pink-900 rounded-full text-white font-bold py-2 px-4 w-full text-sm mb-2">
                                    Abrir
                                </button>
                                {file.Extension === ".svg" &&

                                    <button onClick={(e) => {
                                        copy_clipboard(file, e);
                                    }} className="bg-blue-500 hover:bg-blue-700 rounded-full text-white font-bold py-2 px-4 w-full text-sm">
                                        Copiar
                                    </button>
                                }
                            </div>
                            <div className='border-sky-500 shadow-lime-50 grid grid-cols-1 align-middle items-center justify-items-center'>
                                <div className=''>
                                    <p className="m-0"><a href="#" className='cursor-pointer' onClick={() => {
                                        console.log(file.AbsolutePath);
                                    }}>{file.Name}</a>
                                        <img src={file.Extension === ".svg" ? file.AbsolutePath : "https://cdn3.iconfinder.com/data/icons/file-formats-27/512/Dxf-512.png"} alt="Preview"
                                            className="opacity-100 group-hover:opacity-20
                                            group-hover:w-2/4 m-auto group-hover:m-auto object-contain
                                            object-center overflow-hidden w-64 h-64
                                            duration-500 ease-in-out transition-all bg-slate-300 rounded-lg"
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
