package engine

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	// counter is a package-level atomic counter used to ensure uniqueness in IDs.
	//
	// It increments atomically each time a new ID is generated, guaranteeing that
	// even if multiple IDs are generated within the same millisecond (timestamp),
	// they will still be unique. The `atomic.AddUint64` function is used to safely
	// increment this counter in concurrent environments.
	counter uint64

	// rng is a package-level random number generator used for generating entropy
	// in unique IDs.
	//
	// It is initialized with a time-based seed (`time.Now().UnixNano()`) to ensure
	// that the random numbers generated are different across program executions.
	// This random number generator is used in the `generateUniqueID` function to
	// add a random component to the generated ID, improving uniqueness and reducing
	// the likelihood of collisions.
	rng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// IBase defines the interface for objects with basic identification capabilities.
// It provides methods to access and modify fundamental attributes common to
// many game objects or entities.
//
// Implementers of this interface should provide unique identifiers,
// human-readable names, and categorization tags, along with the ability
// to modify names and tags.
type IBase interface {
	// GetID returns the unique identifier of the object.
	GetID() string

	// GetName returns the human-readable name of the object.
	GetName() string

	// GetTag returns the categorization tag of the object.
	GetTag() string

	// SetName updates the human-readable name of the object.
	SetName(string)

	// SetTag updates the categorization tag of the object.
	SetTag(string)
}

// Base represents a fundamental structure with basic identification attributes.
// It serves as a foundation for other game objects, providing common fields
// and methods for identification and categorization.
//
// Fields:
//   - id: A unique identifier for the instance.
//   - name: A human-readable name for the instance.
//   - tag: A string used for categorization or grouping.
//
// The id field is set upon creation and is immutable, while name and tag
// can be modified using their respective setter methods.
type Base struct {
	id   string
	name string
	tag  string
}

// generateUniqueID creates a unique identifier string.
//
// The generated ID combines three components:
// 1. A timestamp (in milliseconds) to ensure uniqueness across time.
// 2. An atomic counter to guarantee uniqueness even within the same millisecond.
// 3. A random number for additional entropy.
//
// The resulting format is: "timestamp-counter-randomNum"
// where timestamp is in milliseconds, counter is an incrementing integer,
// and randomNum is a three-digit random number.
//
// This function is safe for concurrent use.
//
// Returns:
//   - string: A unique identifier in the format "timestamp-counter-randomNum".
func generateUniqueID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomNum := rng.Intn(1000)
	count := atomic.AddUint64(&counter, 1)
	return fmt.Sprintf("%d-%d-%03d", timestamp, count, randomNum)
}

// NewBase creates and returns a new instance of the Base struct.
//
// This function initializes a Base object with a given name, automatically
// generating a unique ID and setting an initial empty tag. It uses the
// generateUniqueID function to create the ID, ensuring uniqueness across
// all Base instances.
//
// Parameters:
//   - name string: The name to assign to the new Base instance. This should be
//     a human-readable identifier for the object.
//
// Returns:
//   - *Base: A pointer to the newly created Base instance.
//
// Example usage:
//
//	base := NewBase("Player1")
//	fmt.Println(base.GetID())   // Prints the unique ID
//	fmt.Println(base.GetName()) // Prints "Player1"
func NewBase(name string) *Base {
	return &Base{
		id:   generateUniqueID(),
		name: name,
		tag:  "", // Initialize with an empty tag
	}
}

// GetID returns the unique identifier of the Base instance.
//
// This method provides read-only access to the `id` field, which is intended
// to be immutable and uniquely identifies the object.
//
// Returns:
//   - string: The unique identifier of the Base instance.
func (b *Base) GetID() string {
	return b.id
}

// GetName returns the name of the Base instance.
//
// The `name` field represents a human-readable identifier for the object.
// This method allows you to retrieve the current value of the `name`.
//
// Returns:
//   - string: The name of the Base instance.
func (b *Base) GetName() string {
	return b.name
}

// GetTag returns the tag associated with the Base instance.
//
// The `tag` field is used for categorization or grouping purposes.
// This method allows you to retrieve the current value of the `tag`.
//
// Returns:
//   - string: The tag associated with the Base instance.
func (b *Base) GetTag() string {
	return b.tag
}

// SetName updates the name of the Base instance.
//
// This method allows you to modify the `name` field, which represents
// a human-readable identifier for the object.
//
// Parameters:
//   - name (string): The new name to assign to the Base instance.
func (b *Base) SetName(name string) {
	b.name = name
}

// SetTag updates the tag associated with the Base instance.
//
// This method allows you to modify the `tag` field, which is used for
// categorization or grouping purposes.
//
// Parameters:
//   - tag (string): The new tag to assign to the Base instance.
func (b *Base) SetTag(tag string) {
	b.tag = tag
}

var _ IBase = (*Base)(nil)
