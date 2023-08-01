export default function TextInput (props: {
    value: string
    onChange: Function
    placeHolder: string
    label: string
    slim?: boolean
    fullWidth?: boolean
    type?: 'text' | 'email'
    required?: boolean
}) {

    return <label>
        {!!props.label && <div className={"text-sm font-bold mb-1"}>{props.label}</div>}
        <input
        type={props.type ?? "text"}
        className={`bg-gray-100 enabled:text-gray-700 disabled:text-gray-400 ${props.slim ? "py-1" : "py-3"} px-8 rounded disabled:cursor-not-allowed relative ${props.fullWidth ? "w-full" : ""}`}
        placeholder={props.placeHolder}
        value={props.value}
        onChange={(e) => props.onChange(e.target.value)}
        required={props.required === true}
    /></label>
}