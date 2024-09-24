# Dessign-Patterns
Patrones de Diseño en Go
Introducción a los Patrones de Diseño
Los patrones de diseño son soluciones reutilizables a problemas comunes que surgen en el desarrollo de software. Actúan como plantillas que se pueden aplicar en diversas situaciones para abordar problemas de diseño de manera eficiente y eficaz. Estos patrones son el resultado de la experiencia acumulada por desarrolladores a lo largo del tiempo y ayudan a crear código más limpio, organizado y mantenible.

Tipos de Patrones de Diseño
Los patrones de diseño se clasifican generalmente en tres categorías principales:

Patrones Creacionales: Se centran en la forma en que se crean los objetos. Ejemplos incluyen:

Singleton
Factory Method
Abstract Factory
Patrones Estructurales: Se ocupan de la composición de clases y objetos. Ayudan a garantizar que si un sistema se debe expandir, no haya que alterar el código existente. Ejemplos incluyen:

Adapter
Composite
Decorator
Patrones de Comportamiento: Se centran en cómo los objetos interactúan y se comunican entre sí. Ejemplos incluyen:

Observer
Strategy
Command
Patrones Implementados
1. Singleton (Ejemplo: CreateNewAdmin)
Propósito: Asegura que una clase tenga una única instancia y proporciona un punto de acceso global a esa instancia.
Analogía: Piensa en un presidente de un país. Hay un solo presidente al mismo tiempo, y todos los ciudadanos deben dirigirse a esa persona para cuestiones de gobierno.
Reconocimiento: Identificas un patrón Singleton cuando ves que una clase tiene un método que permite obtener una instancia de sí misma y el constructor está privado.
2. Factory Method (Ejemplo: CrearTransporte)
Propósito: Define una interfaz para crear un objeto, pero permite a las subclases alterar el tipo de objeto que se crea.
Analogía: Imagínate una fábrica de coches que puede producir diferentes modelos de automóviles. En lugar de especificar cada tipo de automóvil, se le dice a la fábrica qué tipo de automóvil crear.
Reconocimiento: Se reconoce cuando ves un método que devuelve una instancia de una clase sin especificar el tipo exacto de clase que se va a crear.
3. Observer (Ejemplo: YouTubeChannel)
Propósito: Define una dependencia de uno a muchos entre objetos de tal forma que cuando uno cambia de estado, todos sus dependientes son notificados y actualizados automáticamente.
Analogía: Piensa en un canal de YouTube donde los suscriptores reciben notificaciones cada vez que se publica un nuevo video. Si te suscribes, recibes una alerta.
Reconocimiento: Lo identificas cuando tienes un objeto (el sujeto) que mantiene una lista de suscriptores (observadores) que necesitan ser notificados de los cambios de estado.
4. Strategy (Ejemplo: PaymentProcessor)
Propósito: Define una familia de algoritmos, encapsula cada uno y los hace intercambiables. Permite que el algoritmo varíe independientemente de los clientes que lo utilizan.
Analogía: Imagina que estás en una tienda y puedes elegir entre varios métodos de pago: tarjeta de crédito, PayPal o criptomonedas. Cada método de pago tiene su propio proceso, pero todos logran el mismo objetivo: completar la compra.
Reconocimiento: Se reconoce por la existencia de una interfaz que define una operación y varias implementaciones concretas que la cumplen. Se usa un contexto para llamar al método de la estrategia seleccionada.
Conclusión
Los patrones de diseño son herramientas poderosas para resolver problemas comunes en el desarrollo de software. Entender y aplicar estos patrones puede llevar a un código más limpio, fácil de mantener y flexible. Al usar analogías y reconocer los patrones en ejemplos prácticos, puedes aplicar estos conceptos de manera efectiva en tus proyectos de programación.

Instrucciones Adicionales
Ejemplos de Código: Asegúrate de incluir el código correspondiente a cada patrón en el proyecto para que los lectores puedan ver la implementación en acción.

Práctica: Invita a los lectores a experimentar con los patrones, modificarlos y ver cómo se comportan en diferentes situaciones.

Quitar Comentarios de las Funciones main: Para ejecutar los ejemplos de código, es posible que desees quitar los comentarios de las funciones main. Aquí tienes instrucciones sobre cómo hacerlo en diferentes sistemas operativos:

Windows:

Abre el archivo en un editor de texto como Notepad++ o Visual Studio Code.
Usa la función de búsqueda (Ctrl + F) para encontrar // func main().
Elimina las líneas que comienzan con // y asegúrate de que el bloque de la función main esté activo.
Guarda el archivo.
macOS:

Abre el archivo en un editor de texto como TextEdit o Visual Studio Code.
Usa Cmd + F para buscar // func main().
Elimina las líneas comentadas que comienzan con //.
Guarda los cambios.
Linux:

Abre el archivo en un editor de texto como Gedit o Vim.
Usa Ctrl + F para buscar // func main().
Elimina las líneas comentadas que comienzan con //.
Guarda el archivo.
Una vez que hayas realizado estos cambios, podrás ejecutar los ejemplos de código directamente desde el entorno de Go en tu sistema operativo.

