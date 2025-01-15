import { useState, useEffect } from "react";
import { TextInput, StyleSheet } from "react-native";
import { useThemeColor } from "../hooks/useThemeColor";

interface DebouncedInputProps {
  lightColor?: string;
  darkColor?: string;
  value: string;
  onChange: (value: string) => void;
  debounce?: number;
}

const DebouncedInput: React.FC<DebouncedInputProps> = ({
  lightColor,
  darkColor,
  value: initialValue,
  onChange,
  debounce = 250,
  ...rest
}) => {
  const color = useThemeColor({ light: lightColor, dark: darkColor }, "text");

  const [value, setValue] = useState(initialValue);

  useEffect(() => {
    setValue(initialValue);
  }, [initialValue]);

  useEffect(() => {
    const timeout = setTimeout(() => {
      onChange(value);
    }, debounce);

    return () => clearTimeout(timeout);
  }, [value]);

  return (
    <TextInput
      style={[{ color }, styles.textInput]}
      value={value}
      onChangeText={setValue}
    />
  );
};

const styles = StyleSheet.create({
  textInput: {
    borderWidth: 1,
    borderColor: "#ebebeb",
    padding: 10,
    marginVertical: 5,
  },
});

export default DebouncedInput;
