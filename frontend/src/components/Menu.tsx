export type MenuProps = {
    onsearch: (search: string) => void;
    onselect: (filetype: string) => void;
}
const Menu = (props: MenuProps) => {

    return <>
        <header className="bg-white border-b-sky-800 border shadow shadow-gray-600">
            <nav className="mx-auto flex max-w-7xl lg:max-w-full items-center justify-between p-6 lg:px-8" aria-label="Global">
                <div className="flex lg:flex-1 align-middle items-center">
                    <a href="#" className="text-2xl text-gray-950 font-semibold">
                        Darkburn
                    </a>

                    <div className='w-2/4 m-auto ml-4'>
                        <input type="text" placeholder='Buscar' className="border rounded-full border-gray-500 w-full p-2 text-black" onChange={(e) => {
                            console.log(e.target.value);
                            props.onsearch(e.target.value);
                        }} />
                    </div>

                    <div className='w-2/4 m-auto'>
                        <select className="border rounded-full ml-2 border-gray-500 w-full p-1 text-black" onChange={(e) => {
                            console.log(e.target.value);
                            props.onselect(e.target.value);
                        }}>
                            <option value="todos">Todo</option>
                            <option value=".svg">SVG</option>
                            <option value=".dxf">DXF</option>
                        </select>
                    </div>
                </div>
                <div className="flex lg:hidden">
                    <button type="button" className="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700">
                        <span className="sr-only">Open main menu</span>
                        <svg className="size-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true" data-slot="icon">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                        </svg>
                    </button>
                </div>
                <div className="hidden lg:flex lg:flex-1 lg:justify-end">
                    <a href="#" className="text-sm/6 font-semibold text-gray-900">AATS91</a>
                </div>
            </nav>
        </header>
    </>
}
export default Menu;
