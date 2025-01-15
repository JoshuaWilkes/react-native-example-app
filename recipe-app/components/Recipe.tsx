import { Recipe as RecipeType } from "../gen/recipe/v1alpha1/recipe_pb";
import { Collapsible } from "./Collapsible";
import { ThemedText } from "./ThemedText";
import { ThemedView } from "./ThemedView";

type RecipeProps = { recipe: RecipeType };

const Recipe: React.FC<RecipeProps> = ({ recipe }) => {
  return (
    <ThemedView>
      <ThemedText type="subtitle">{recipe.title}</ThemedText>
      <Collapsible title="Ingredients">
        {recipe.ingredients.map((ingredient, j) => {
          return (
            <ThemedText key={`rec-${recipe.id}-${j}`}>{ingredient}</ThemedText>
          );
        })}
      </Collapsible>
      <Collapsible title="Recipe">
        <ThemedText>{recipe.body}</ThemedText>
      </Collapsible>
    </ThemedView>
  );
};

export default Recipe;
