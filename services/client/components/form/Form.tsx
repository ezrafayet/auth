export default function Form (props: {
    onSubmit: any,
    children: any
}) {
    return <form className={"w-95p max-w-form rounded-xl py-5 px-3 bg-white md:px-9 flex flex-col items-center shadow-gray-500 shadow-md"}
                 onSubmit={props.onSubmit}>
        {props.children}
    </form>
}