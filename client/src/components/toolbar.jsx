const Toolbar = () => {
  return (
    <div className="flex h-14 items-center bg-stone-50 shadow-md">
      <img
        src="/images/pixel-art-joystick.png"
        alt="Joystick"
        id="joystick-icon"
        className={`ml-7 h-9 w-8 cursor-pointer transition-all duration-150 active:scale-110`}
      />
      <div className="ml-40 flex h-5 items-center">
        <ToolbarLink content={"PRODUCTS"} showLine={true} path={"/"} />
        <ToolbarLink content={"COMMUNITY"} showLine={true} path={"/"} />
        <ToolbarLink content={"GROUP BUYS"} showLine={true} path={"/"} />
        <ToolbarLink content={"EVENTS"} showLine={false} path={"/"} />
      </div>
      <div className="flex h-full w-full items-center justify-end">
        <button className="mr-3 flex h-8 w-20 items-center justify-center rounded-sm bg-indigo-600 hover:scale-105 hover:opacity-95">
          <p className="font-nunito text-xs font-medium tracking-wide text-white">
            LOGIN
          </p>
        </button>
        <button className="mr-3 flex h-8 w-20 items-center justify-center rounded-sm bg-white hover:scale-105 hover:bg-stone-200">
          <p className="font-nunito text-xs font-medium tracking-wide">LOGIN</p>
        </button>
      </div>
    </div>
  );
};

const ToolbarLink = ({ content, showLine, path }) => {
  return (
    <div className="flex h-full items-center">
      <p className="cursor-pointer select-none whitespace-nowrap font-nunito text-xs font-medium tracking-widest text-neutral-700">
        {content}
      </p>
      <div
        className={`ml-2 mr-2 h-full w-[1px] rounded-md bg-neutral-700 opacity-50 ${
          showLine ? "" : "hidden"
        }`}
      />
    </div>
  );
};

export default Toolbar;
