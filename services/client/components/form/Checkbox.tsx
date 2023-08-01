export default function Checkbox (props: {
    text: string
    checked: boolean
    onChange: Function
    disabled?: boolean
    required?: boolean
}) {

    return <label className={"flex items-center"}>
            <input required={props.required === true} type={"checkbox"} checked={props.checked} onChange={() => props.onChange()} disabled={props.disabled} className={"mr-2 w-5 h-5 shrink-0"} />
            <div className={"text-sm"}>{props.text}</div>
        </label>
}