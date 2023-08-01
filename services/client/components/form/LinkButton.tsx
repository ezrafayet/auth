
export default function LinkButton (props: {
    text: string
    onClick?: () => void
    disabled?: boolean
    infobox?: string
}) {

    return <button disabled={props.disabled}
                   onClick={() => props.onClick ? props.onClick() : null}
                   title={props.infobox}
                   type={"button"}
                   className={"text-blue-600 hover:text-blue-500 transition-colors delay-200"}>
        {props.text}
    </button>
}