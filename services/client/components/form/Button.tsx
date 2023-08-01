
export default function (props: {
    text: string
    onClick?: () => void
    disabled?: boolean
    fullWidth?: boolean
    infobox?: string
    loading?: boolean
    slim?: boolean
    submit?: boolean
}) {

    return <button
        className={`bg-gray-300 enabled:hover:bg-gray-400 enabled:text-gray-700 disabled:text-gray-400 text-sm font-bold ${props.slim ? "py-1" : "py-3"} px-8 rounded transition-colors duration-200 enabled:cursor-pointer disabled:cursor-not-allowed relative ${props.fullWidth ? "w-full" : ""}`}
        onClick={() => props.onClick ? props.onClick() : null}
        title={props.infobox}
        disabled={props.disabled}
        type={props.submit ? "submit" : "button"}
    >
        { props.text }
        { props.loading && <img className={"w-4 h-4 absolute bottom-1 right-1"} src={"/spinner.svg"} /> }
    </button>
}